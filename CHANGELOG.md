# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

- Rebooted to simplify structure and get some basics right before too deep.
- Following [a tutorial](https://kb.objectrocket.com/category/postgresql?filter=Go+Lang+and+PostgreSQL+Web+App+MVC+pattern+Part).
  - Part 2 completed.
- Table migration configured.
- Basic user model and seeder created.
- Save records with a hashed password

### ERRORS

1. The password should be encrypted. So...
1. Need to have show/edit for user.
1. Need to have an authorization so (for example) the users list doesn't display.
1. Create a separate folder for each of the data elements? eg `models/user/` `models/songs` and so on?
1. To create independant homepages with user logged in. The cookie should have the user id embedded.

## v0.0.1 2020-10-17

- Basic landing page added.
- Load database setup from environment.
- Initial migration for user configured.
