# pacchetto

## What is it?
A Go library and application that automates the creation of Assetto Corsa
dedicated server packages.

## What does that mean?
That you can create an archive and upload it to your Assetto Corsa host.
This means that you do not have to download and sign into Steam on the host.
It also allows you to store clean copies of the game data somewhere.

## How do I use it?
The application can create two types of packages:

* A single `phat` package file
* Several smaller `distributed` packages

Either of these formats can be used. However, you may prefer using one over
another depending on the situation. For example, if you have an unstable
internet connection, you may prefer using the `distributed` packaging mode to
mitigate the chances of an upload failure.

#### Creating a phat package
Execute the following on the command line:
```
pacchetto -m phat
```

You can also run:
```
pacchetto -m p
```

#### Creating distributed packages
Execute the following on the command line:
```
pacchetto -m distributed
```

You can also run:
```
pacchetto -m d
```

## How do I build it?
Please refer to the [building documentation](docs/building/README.md).