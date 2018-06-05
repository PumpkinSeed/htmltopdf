#!/bin/sh

VERSION="0.12.4"

wget https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/$VERSION/wkhtmltox-0.12.4_linux-generic-amd64.tar.xz

tar -xf wkhtmltox-0.12.4_linux-generic-amd64.tar.xz

cp wkhtmltox/bin/wkhtmltoimage /usr/bin/
cp wkhtmltox/bin/wkhtmltopdf /usr/bin/

ls /usr/bin/ | grep wkhtmlto

echo $PATH

wkhtmltopdf http://google.com google.pdf 