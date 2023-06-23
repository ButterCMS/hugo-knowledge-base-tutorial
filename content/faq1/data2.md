---
title: Install Hugo on Linux 
description: "Getting started with Hugo"
date: 2023-04-12
---

## Install Hugo on Linux
##### Date: 2023-04-12



Choose How to Install
If you want to use Hugo as your site generator, simply install the Hugo binaries.

To contribute to the Hugo source code or documentation, you should fork the Hugo GitHub project and clone it to your local machine.

Finally, you can install the Hugo source code with go, build the binaries yourself, and run Hugo that way. Building the binaries is an easy task for an experienced go getter.

Install Hugo as Your Site Generator (Binary Install)
Use the installation instructions in the Hugo documentation.

Build and Install the Binary from Source (Using the Go toolchain)
Prerequisite Tools
Go (we test it with the last 2 major versions; but note that Hugo 0.95.0 only builds with >= Go 1.18.)
Fetch from GitHub
To fetch, build and install from the Github source:

go install github.com/gohugoio/hugo@latest
If you want to compile with Sass/SCSS support use -tags extended and make sure CGO_ENABLED=1 is set in your go environment. If you don't want to have CGO enabled, you may use the following command to temporarily enable CGO only for hugo compilation:

CGO_ENABLED=1 go install -tags extended github.com/gohugoio/hugo@latest
The Hugo Documentation
The Hugo documentation now lives in its own repository, see https://github.com/gohugoio/hugoDocs. But we do keep a version of that documentation as a git subtree in this repository. To build the sub folder /docs as a Hugo site, you need to clone this repo:

git clone git@github.com:gohugoio/hugo.git
Contributing code to Hugo
For a complete guide to contributing to Hugo, see the Contribution Guide.

We welcome contributions to Hugo of any kind including documentation, themes, organization, tutorials, blog posts, bug reports, issues, feature requests, feature implementations, pull requests, answering questions on the forum, helping to manage issues, etc.

The Hugo community and maintainers are very active and helpful, and the project benefits greatly from this activity.

Asking Support Questions
We have an active discussion forum where users and developers can ask questions. Please don't use the GitHub issue tracker to ask questions.

Reporting Issues
If you believe you have found a defect in Hugo or its documentation, use the GitHub issue tracker to report the problem to the Hugo maintainers. If you're not sure if it's a bug or not, start by asking in the discussion forum. When reporting the issue, please provide the version of Hugo in use (hugo version).

Dependencies
Hugo stands on the shoulder of many great open source libraries.

If you run hugo env -v you will get a complete and up to date list.

In Hugo 0.111.2 that list is, in lexical order:

page_type: hugo_knowledge_base

published: 2023-04-12

slug: how-to-install-hugo-on-windows

updated: 2023-04-12T19:33:36.893950Z

