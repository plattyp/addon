# Heroku Addon Example [![Code Climate](https://codeclimate.com/github/plattyp/addon/badges/gpa.svg)](https://codeclimate.com/github/plattyp/addon) [![Build Status](https://travis-ci.org/plattyp/addon.svg?branch=add-travis)](https://travis-ci.org/plattyp/addon)

An example Heroku addon written in Golang

## Why use this?

Heroku gives a ton of documentation throughout their site, but I couldn't find a good end-to-end example of creating a custom Heroku addon that can then be published into their Marketplace. I wanted this to cover all parts of the addon from the provisioning and SSO endpoints to the actual table creation and database model.

## Assumptions

  - Using a PostgreSQL database as a datastore (Install Postgres locally)
  - The ENVs in this are just used as an example, real world you'd probably use a vault or store them only on the deployed environment

## Structure

There are 3 tables created as part of this: `plans`, `users`, and `apps`. The assumption is that a User can have many apps or in Heroku terms this would be a single Heroku user could SSO once into this addon application and have access to their various Heroku associated apps.

## Setup dependencies

    brew install glide

## Creating DB / Running Migrations & Seeds

    createdb addon
    sql-migrate up

## Building It

    make

## Running It (exposed on port 5000 by default)

    make run

## Test It

This is a complete example. I still have some additional tests to right, but it currently passes the Addon spec. This has been tested against the [kensa gem](https://github.com/heroku/kensa). Travis executes the following commands after inserting the test data.

    kensa test
    kensa test sso 123
