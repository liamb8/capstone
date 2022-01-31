

# VSFTPd Lab

**Virtual Machine Download Link**

------

Before starting the lab, make sure to **download the ZIP file provided**. After downloading, **unzip the file** and place the folder wherever you find appropriate. The instructions below assume that you're using VMware Workstation Client, although you can still follow them if you're using other VM clients.

Go to the **File** tab in the upper left-hand corner of the VM clients window and click on it, a drop-down menu will show up. From there, click on **Open**. After that, a file explorer will open navigate to where you put the VM you unzipped. Click on the `SSHHardeningLab.vmx` image and it should show up in your client. If this doesn't work click on the `SSHHardeningLab.ovf` and it will bring up an import menu. 

Once you have done this click on the **Import** button and the VM should show up. After this right click on the VM and click on **settings**, this is where you can change the settings for your VM. Below is an image of what the settings of the VM should be. After checking that all the settings are the same click **Ok** and then right-click on the VM and choose power and then click on **Start up guest** to power on the VM.


### User Accounts:

Below is the user account which will be used for this lab.

- John - Administrator
  - User Login: `john`
  - Password: `Fallentrees4`

------



## Setup VSFTPd

### Change into the root user

```
sudo su
```

### Install VSFTPd

```
apt-get install vsftpd -y
```

### Start VSFTPd and set it to start on boot

```
systemctl start vsftpd
systemctl enable vsftpd
```

### Create a user for FTP access

```
adduser vsftp
```

### Make an FTP directory and set permissions

```
mkdir /home/vsftp/ftp

chown nobody:nogroup /home/vsftp/ftp

chmod a-w /home/vsftp/ftp
```

### Create an upload directory and set permissions

```
mkdir /home/vsftp/ftp/test

chown vsftp:vsftp /home/vsftp/ftp/test
```

------

## Configure VSFTPd

### Backup the configuration file

```
cp /etc/vsftpd.conf /etc/vsftpd.conf.bak
```

### Open the configuration file in your favourite text editor

```
vi /etc/vsftpd.conf
```

### Add the following lines to the file, then save and close the file:

```
listen=NO
listen_ipv6=YES
anonymous_enable=NO
local_enable=YES
write_enable=YES
local_umask=022
dirmessage_enable=YES
use_localtime=YES
xferlog_enable=YES
connect_from_port_20=YES
chroot_local_user=YES
secure_chroot_dir=/var/run/vsftpd/empty
pam_service_name=vsftpd
pasv_enable=Yes
pasv_min_port=10000
pasv_max_port=11000
user_sub_token=$USER
local_root=/home/$USER/ftp
userlist_enable=YES
userlist_file=/etc/vsftpd.userlist
userlist_deny=NO
rsa_cert_file=/etc/cert/vsftpd.pem
rsa_private_key_file=/etc/cert/vsftpd.pem
ssl_enable=YES
allow_anon_ssl=NO
force_local_data_ssl=YES
force_local_logins_ssl=YES
ssl_tlsv1=YES
ssl_sslv2=NO
ssl_sslv3=NO
require_ssl_reuse=NO
ssl_ciphers=HIGH
```

### Add the FTP user to VSFTP

```
vi /etc/vsftpd.userlist
```

Add the following line, then save and close the file:

```
vsftp
```

### Create a certificate to connect via SSL

```
mkdir /etc/cert

openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /etc/cert/vsftpd.pem -out /etc/cert/vsftpd.pem
```

### Restart VSFTP

```
systemctl restart vsftpd
```

------

## How to Connect to FTP Server:

Visit ftp://YOUR-SERVER-IP and login using the username and password created earlier in order to view files uploaded.

------

## Secure VSFTPd:

- edit /etc/vsftpd.conf file by doing `nano /etc/vsftpd.conf`
- Uncomment **ftpd_banner** and change the text to anything you want or leave it blank to hide the version. Then press `Cntrl + X` to save the file.

- Disable anonymous login for FTP server by editing the FTP configuration file `/etc/vsftpd.conf` and commenting out anonymous_enable line or change the attribute to no from yes.

- SSL/TLS should be implemented to ensure the communication is encrypted between server and client and the attacker cannot read the ciphertext. This has already been setup in the previous steps by creating the SSL certificate and editing the `/etc/vsftpd.conf` file to use SSL. 