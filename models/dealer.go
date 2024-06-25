package models

type Dealer interface {
	Deal() (Card, error)
}
