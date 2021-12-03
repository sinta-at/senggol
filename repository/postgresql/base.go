package postgresql

import (
	"database/sql"
	"fmt"
	"senggol/repository"
	"time"

	_ "github.com/lib/pq"
)

type PostgresqlRepositories struct {
	CreateUser                repository.CreateUser
	GetUserByUsername         repository.GetUserByUsername
	GetPeers                  repository.GetPeers
	GetPeersCount             repository.GetPeersCount
	GetPeerDirectMessages     repository.GetPeerDirectMessages
	CreateDirectMessage       repository.CreateDirectMessage
	UpdateDirectMessageSeenAt repository.UpdateDirectMessageSeenAt
}

func GetRepositories(host, port, user, dbname, password, sslmode string) (PostgresqlRepositories, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode))
	if err != nil {
		return PostgresqlRepositories{}, err
	}

	maxLifeTime := 10 * time.Second
	maxIdle, maxOpenConnection := 10, 10
	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxOpenConnection)
	db.SetConnMaxLifetime(maxLifeTime)

	return PostgresqlRepositories{
		CreateUser:                CreateUserRepository{db},
		GetUserByUsername:         GetUserByUsernameRepository{db},
	    GetPeers:                  GetPeersRepository{db},
	    GetPeersCount:             GetPeersCountRepository{db},
	    GetPeerDirectMessages:     GetPeerDirectMessagesRepository{db},
	    CreateDirectMessage:       CreateDirectMessageRepository{db},
	    UpdateDirectMessageSeenAt: UpdateDirectMessageSeenAtRepository{db},
	}, nil
}
