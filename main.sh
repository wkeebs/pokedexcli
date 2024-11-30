#!/bin/zsh
rm build/*
go build -C src -o ../build/pokedexcli 
./build/pokedexcli