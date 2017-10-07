# Addon (WIP) [![Code Climate](https://codeclimate.com/github/plattyp/addon/badges/gpa.svg)](https://codeclimate.com/github/plattyp/addon) [![Build Status](https://travis-ci.org/plattyp/addon.svg?branch=add-travis)](https://travis-ci.org/plattyp/addon)

An example Heroku addon written in Golang

## Why use this?

Heroku gives a ton of documentation throughout their site, but I couldn't find a good end-to-end example of creating a custom Heroku addon that can then be published into their Marketplace. I wanted this to cover all parts of the addon from the provisioning and SSO endpoints to the actual table creation and database model.

## Assumptions

  - Using a PostgreSQL database as a datastore (Install Postgres locally)

## Setup dependencies

    brew install glide

## Creating DB / Running Migrations & Seeds

    createdb addon
    sql-migrate up

## Building It

    make

## Running It

    make run

## Current Status

This is a WIP. Will update with specs and more installation instructions once it is complete.
