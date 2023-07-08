# Implementation of a Sequential list backed by an array in Go using TDD


## v0: implementation in package main

In this version we use TDD to build the implementation of the list data structure backed by an array of a certain fixed number of elements (ten).

## v1: creating a list package

In this version, we create a `list` package and move the implementation there.

Also a few example usages are illustrated in `main.go`.

## v2: list package with generics

In this step we introduce generics, so that the list can work with any type.

## ToDo

Make the list generic so that it can work with any type.