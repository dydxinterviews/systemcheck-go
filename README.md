<p align='center'><img src='https://dydx.exchange/logo.svg' width='256' /></p>

The goal of this repository is to verify that your development environment is set up to start working on code at dYdX. Successfully running the tests in this repo confirms that your environment is close to or completely set up to work with dYdX repositories.

## SSH Key

Setup your ssh key be able to clone the repository by following the instructions [here](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account?platform=mac)

## Git
Clone the repository with one of the following
```
git clone git@github.com:dydxprotocol/go-systemCheck.git
```

Create a new branch locally

`git checkout -b <branch_name>`

Stage files to be committed

`git add <file_names>`

Add a commit with a message

`git commit -m "<message>"`

Push to the remote branch

`git push origin <branch_name>`

Create the pull request on github.com from your branch to master

[Here](http://guides.beanstalkapp.com/version-control/common-git-commands.html) is a list of helpful git commands.

## Editor

Pick a text editor you are comfortable with. Most editors should have javascript support. Many engineers at dYdX use [VSCode](https://code.visualstudio.com) or [Atom](https://atom.io).

### Install

1. Install Go `1.18` or newer [here](https://go.dev/doc/install).
2. Run `make install` to install project dependencies.

### Build

After installing the dependencies, you can build the service using:

```
make build
```

### Tests

```
make test
```

Note: most of our test assertins are powered by [testify](https://github.com/stretchr/testify).
