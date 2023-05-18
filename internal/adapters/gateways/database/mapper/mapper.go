package mapper

import "github.com/victorskg/my-wallet/pkg/database"

type EntityToDomainMapper[T database.Entity, P any] interface {
	FromEntityToDomain(entity T) *P
}

type DomainToEntityMapper[T database.Entity, P any] interface {
	FromDomainToEntity(domain P) *T
}

type Mapper[T database.Entity, P any] interface {
	FromEntityToDomain(entity T) *P
	FromDomainToEntity(domain P) *T
}
