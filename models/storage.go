package models

import "context"

type Storage interface {
	GetIPs(context.Context) ([]IP, error)
	AddIP(context.Context, IP) (IP, error)
	RemoveIP(context.Context, string) error
}
