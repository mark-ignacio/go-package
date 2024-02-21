# go-package

**go-package** builds Linux packages for Go binaries.

The package is an automated version of the 3 step process in [https://go.dev/doc/install](https://go.dev/doc/install).

## Repositories

Releases are uploaded to [Gemfury](https://gemfury.com) hosted package repositories @ https://yum.fury.io/iio/ and https://apt.fury.io/iio/.

Usage instructions below:

### yum

<!-- sudo rpm --import https://yum.fury.io/iio/gpg.key -->

```sh
echo '[fury-iio]
name=iio
baseurl=https://yum.fury.io/iio/
enabled=1
gpgcheck=0' | sudo tee /etc/yum.repos.d/fury-iio.repo
sudo dnf install golang
```

### apt

<!-- curl https://apt.fury.io/iio/gpg.key | sudo apt-key add -->

```sh
echo "deb https://apt.fury.io/iio/ * *" | sudo tee /etc/apt/sources.list.d/fury-iio.list
sudo apt update
sudo apt install golang
```
