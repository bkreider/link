package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/Scalingo/link/api"
	"github.com/Scalingo/link/models"
	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli"
)

var Version = "dev"

func main() {
	app := cli.NewApp()

	app.Version = Version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "127.0.0.1:1313",
			Usage: "Host to contact",
		},
		cli.StringFlag{
			Name:  "user, u",
			Value: "",
			Usage: "Username for basic auth",
		},
		cli.StringFlag{
			Name:  "password, p",
			Value: "",
			Usage: "Password for basic auth",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "list",
			Action: func(c *cli.Context) error {
				client := getClientFromCtx(c)
				ips, err := client.ListIPs(context.Background())
				if err != nil {
					return err
				}
				formatIPs(ips)
				return nil
			},
		}, {
			Name:      "destroy",
			ArgsUsage: "ID",
			Action: func(c *cli.Context) error {
				if c.NArg() != 1 {
					cli.ShowCommandHelp(c, c.Command.Name)
					return nil
				}
				client := getClientFromCtx(c)
				err := client.RemoveIP(context.Background(), c.Args().First())
				if err != nil {
					return err
				}
				fmt.Println(aurora.Green(fmt.Sprintf("IP %v deleted.", c.Args().First())))
				return nil
			},
		}, {
			Name:      "get",
			ArgsUsage: "ID",
			Action: func(c *cli.Context) error {
				if c.NArg() != 1 {
					cli.ShowCommandHelp(c, c.Command.Name)
					return nil
				}
				client := getClientFromCtx(c)
				ip, err := client.GetIP(context.Background(), c.Args().First())
				if err != nil {
					return err
				}
				formatIP(ip)
				return nil
			},
		}, {
			Name:      "update",
			ArgsUsage: "ID",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "check-type", Value: string(api.TCPHealthCheck), Usage: "Type of healthcheck"},
				cli.StringFlag{Name: "check-host", Value: "", Usage: "Host for healthcheck"},
				cli.IntFlag{Name: "check-port", Value: 0, Usage: "Port for healthcheck"},
			},
			Action: func(c *cli.Context) error {
				if c.NArg() != 1 {
					cli.ShowCommandHelp(c, c.Command.Name)
					return nil
				}
				if c.String("check-host") == "" || c.Int("check-port") == 0 {
					cli.ShowCommandHelp(c, c.Command.Name)
					return nil
				}
				opts := api.UpdateIPOpts{
					Checks: []models.Healthcheck{{
						Type: models.HealthcheckType(c.String("check-type")),
						Host: c.String("check-host"),
						Port: c.Int("check-port"),
					}},
				}
				client := getClientFromCtx(c)
				ip, err := client.UpdateIP(context.Background(), c.Args().First(), opts)
				if err != nil {
					return err
				}
				formatIP(ip)
				return nil
			},
		}, {
			Name:      "try-get-lock",
			ArgsUsage: "ID",
			Action: func(c *cli.Context) error {
				if c.NArg() != 1 {
					cli.ShowCommandHelp(c, c.Command.Name)
					return nil
				}
				client := getClientFromCtx(c)
				err := client.TryGetLock(context.Background(), c.Args().First())
				if err != nil {
					return err
				}
				fmt.Println(aurora.Green("Request sent."))
				return nil
			},
		}, {
			Name:      "add",
			ArgsUsage: "IP [CHECK_TYPE CHECK_ENDPOINT]...",
			Action: func(c *cli.Context) error {
				if len(c.Args())%2 == 0 {
					// 1 For the IP
					// And 2 per Healthchecks
					// So NArgs % 2 must be == 1
					cli.ShowCommandHelp(c, c.Command.Name)
					return nil
				}
				client := getClientFromCtx(c)
				ip := c.Args().First()
				var checks []models.Healthcheck
				curArg := 1
				for curArg < c.NArg() {
					endpoint := c.Args().Get(curArg + 1)
					host, port, err := net.SplitHostPort(endpoint)
					if err != nil {
						return fmt.Errorf("Invalid endpoint: %s", endpoint)
					}
					portI, err := strconv.Atoi(port)
					if err != nil {
						return fmt.Errorf("Invalid endpoint: %s", endpoint)
					}
					checks = append(checks, models.Healthcheck{
						Type: models.HealthcheckType(c.Args().Get(curArg)),
						Host: host,
						Port: portI,
					})
					curArg += 2
				}

				newIP, err := client.AddIP(context.Background(), ip, checks...)
				if err != nil {
					return err
				}

				fmt.Println(aurora.Green(fmt.Sprintf("IP %s (%s) successfully added", newIP.IP.IP, newIP.ID)))
				return nil
			},
		}, {
			Name: "version",
			Action: func(c *cli.Context) error {
				fmt.Printf("Client version: \t%s\n", app.Version)
				client := getClientFromCtx(c)
				version, err := client.Version(context.Background())
				if err != nil {
					version = aurora.Red(err.Error()).String()
				}
				fmt.Printf("Server version: \t%s\n", version)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(aurora.Red(fmt.Sprintf("Error: %s", err.Error())))
	}
}

func getClientFromCtx(c *cli.Context) api.HTTPClient {
	var opts []api.ClientOpt
	if c.GlobalString("host") != "" {
		opts = append(opts, api.WithURL(c.GlobalString("host")))
	}

	if c.GlobalString("user") != "" {
		opts = append(opts, api.WithUser(c.GlobalString("user")))
	}

	if c.GlobalString("password") != "" {
		opts = append(opts, api.WithPassword(c.GlobalString("password")))
	}

	return api.NewHTTPClient(opts...)
}
