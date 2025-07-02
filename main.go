package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Dass33/gator/internal/config"
	"github.com/Dass33/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	dbQueries := database.New(db)

	st := state{
		config: &cfg,
		db:     dbQueries,
	}

	cmds := commands{
		data: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Not enough arguments provided")
		os.Exit(1)
	}
	com := command{
		name:      args[1],
		arguments: args[2:],
	}
	err = cmds.run(&st, com)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
