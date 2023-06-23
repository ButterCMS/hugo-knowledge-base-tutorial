---
title: Install Hugo on macOS
description: "Issues with installation"
date: 2023-04-12
---

## Install Hugo on macOS
##### Date: 2023-04-12




This tutorial aims to be a complete guide to installing Hugo on your Mac computer.

Assumptions
You know how to open a terminal window.
You’re running a modern 64-bit Mac.
You will use ~/Sites as the starting point for your site.
Pick Your Method
There are three ways to install Hugo on your Mac computer: the brew utility, from the distribution, or from source. There’s no “best” way to do this. You should use the method that works best for your use case.

There are pros and cons for each.

Brew is the simplest and least work to maintain. The drawbacks aren’t severe. The default package will be for the most recent release, so it will not have bug-fixes until the next release (unless you install it with the --HEAD option). The release to brew may lag a few days behind because it has to be coordinated with another team. Still, I’d recommend brew if you want to work from a stable, widely used source. It works well and is really easy to update.

Downloading the tarball and installing from it is also easy. You have to have a few more command line skills. Updates are easy, too. You just repeat the process with the new binary. This gives you the flexibility to have multiple versions on your computer. If you don’t want to use brew, then the binary is a good choice.

Compiling from source is the most work. The advantage is that you don’t have to wait for a release to add features or bug fixes. The disadvantage is that you need to spend more time managing the setup. It’s not a lot, but it’s more than with the other two options.

Since this is a “beginner” how-to, I’m going to cover the first two options in detail and go over the third more quickly.

Brew
Step 1: Install brew if you haven’t already
Go to the brew website, http://brew.sh/, and follow the directions there. The most important step is:

ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
When I did this, I had some problems with directory permissions. Searches on Google pointed me to pages that walked me through updating permissions on the /usr/local directory. Seemed scary, but it’s worked well since.

Step 2: Run the brew command to install hugo
First, update the formulae and Homebrew itself by running:

$ brew update
Then, install Hugo using Homebrew by running:

$ brew install hugo
==> Downloading https://homebrew.bintray.com/bottles/hugo-0.13_1.yosemite.bottle.tar.gz
######################################################################## 100.0%
==> Pouring hugo-0.13_1.yosemite.bottle.tar.gz
🍺  /usr/local/Cellar/hugo/0.13_1: 4 files,  14M
(Note: Replace brew install hugo with brew install hugo --HEAD if you want the absolute latest version in development, but beware—there might be bugs!)

Brew should have updated your path to include Hugo. Confirm by opening a new terminal window and running a few commands:

$ # show the location of the hugo executable
$ which hugo
/usr/local/bin/hugo

$ # show the installed version
$ ls -l $( which hugo )
lrwxr-xr-x  1 mdhender admin  30 Mar 28 22:19 /usr/local/bin/hugo -> ../Cellar/hugo/0.13_1/bin/hugo

$ # verify that hugo runs correctly
$ hugo version
Hugo Static Site Generator v0.13 BuildDate: 2015-03-09T21:34:47-05:00
Step 3: You’re Done
You’ve installed Hugo. Now you need to set up your site. Read the Quickstart guide, explore the rest of the documentation, and if you still have questions just ask!

From Tarball
Step 1: Decide on the location
When installing from the tarball, you have to decide if you’re going to install the binary in /usr/local/bin or in your home directory. There are three camps on this:

Install it in /usr/local/bin so that all the users on your system have access to it. This is a good idea because it’s a fairly standard place for executables. The downside is that you may need elevated privileges to put software into that location. Also, if there are multiple users on your system, they will all run the same version. Sometimes this can be an issue if you want to try out a new release.

Install it in ~/bin so that only you can execute it. This is a good idea because it’s easy to do, easy to maintain, and doesn’t require elevated privileges. The downside is that only you can run Hugo. If there are other users on your site, they have to maintain their own copies. That can lead to people running different versions. of course, this does make it easier for you to experiment with different releases.

Install it in your sites directory. This is not a bad idea if you have only one site that you’re building. It keeps every thing in a single place. If you want to try out new releases, you can just make a copy of the entire site, update the Hugo executable, and have it.

All three locations will work for you. I’m going to document the second option, mostly because I’m comfortable with it.

Step 2: Download the Tarball
Open https://github.com/spf13/hugo/releases in your browser.

Find the current release by scrolling down and looking for the green tag that reads “Latest Release.”

Download the current tarball for the Mac. The name will be something like hugo_X.Y_osx-64bit.tgz, where X.YY is the release number.

By default, the tarball will be saved to your ~/Downloads directory. If you chose to use a different location, you’ll need to change that in the following steps.

Step 3: Confirm your download
Verify that the tarball wasn’t corrupted during the download:

$ tar tvf ~/Downloads/hugo_X.Y_osx-64bit.tgz
-rwxrwxrwx  0 0      0           0 Feb 22 04:02 hugo_X.Y_osx-64bit/hugo_X.Y_osx-64bit.tgz
-rwxrwxrwx  0 0      0           0 Feb 22 03:24 hugo_X.Y_osx-64bit/README.md
-rwxrwxrwx  0 0      0           0 Jan 30 18:48 hugo_X.Y_osx-64bit/LICENSE.md
The .md files are documentation. The other file is the executable.

Step 4: Install into your bin directory
$ # create the directory if needed
$ mkdir -p ~/bin

$ # make it the working directory
$ cd ~/bin

$ # extract the tarball
$ tar -xvzf ~/Downloads/hugo_X.Y_osx-64bit.tgz
Archive:  hugo_X.Y_osx-64bit.tgz
  x ./
  x ./hugo
  x ./LICENSE.md
  x ./README.md

$ # verify that it runs
$ ./hugo version
Hugo Static Site Generator v0.13 BuildDate: 2015-02-22T04:02:30-06:00
You may need to add your bin directory to your PATH variable. The which command will check for us. If it can find hugo, it will print the full path to it. Otherwise, it will not print anything.

$ # check if hugo is in the path
$ which hugo
/Users/USERNAME/bin/hugo
If hugo is not in your PATH, add it by updating your ~/.bash_profile file. First, start up an editor:

$ nano ~/.bash_profile
Add a line to update your PATH variable:

export PATH=$PATH:$HOME/bin
Then save the file by pressing Control-X, then Y to save the file and return to the prompt.

Close the terminal and then open a new terminal to pick up the changes to your profile. Verify by running the which hugo command again.

Step 5: You’re Done
You’ve installed Hugo. Now you need to set up your site. Read the Quickstart guide, explore the rest of the documentation, and if you still have questions just ask!

Building from Source
If you want to compile Hugo yourself, you’ll need Go, which is also available from Homebrew: brew install go.

Step 1: Get the Source


name: How to install Hugo on my Macbook

page_type: hugo_knowledge_base

published: 2023-04-11

slug: how-to-install-hugo-on-my-macbook

updated: 2023-04-12T17:14:39.277349Z

