# Apache Lab

**Virtual Machine Download Link**

------

Before starting the lab, make sure to **download the ZIP file provided**. After downloading, **unzip the file** and place the folder wherever you find appropriate. The instructions below assume that you're using VMware Workstation Client, although you can still follow them if you're using other VM clients.

Go to the **File** tab in the upper left-hand corner of the VM clients window and click on it, a drop-down menu will show up. From there, click on **Open**. After that, a file explorer will open navigate to where you put the VM you unzipped. Click on the `SSHHardeningLab.vmx` image and it should show up in your client. If this doesn't work click on the `SSHHardeningLab.ovf` and it will bring up an import menu. 

Once you have done this click on the **Import** button and the VM should show up. After this right click on the VM and click on **settings**, this is where you can change the settings for your VM. Below is an image of what the settings of the VM should be. After checking that all the settings are the same click **Ok** and then right-click on the VM and choose power and then click on **Start up guest** to power on the VM.
![](https://github.com/liamb8/capstone/blob/main/Labs/SSH%20Hardening/Pictures/vmsettings.jpg)

### User Accounts:

Below is the user account which will be used for this lab.

- John - Administrator - Ubuntu VM
  - User Login: `john`
  - Password: `Fallentrees4`
- Kali - Administrator - Kali Linux VM
  - User Login: `kali`
  - Password: `Fallentrees4`

------

## Install Apache



## Run the Script



## Hardening Steps



### Apache Version Hiding

In the default configuration for Apache the web server version is exposed and available for anyone to see. You don't want to expose the version of the web server you're using as this can help hackers identify vulnerabilities with that specific version. As we can see with a wireshark capture of just going to the web page it will display the Apache version info.

![image-20220210191444331](C:\Users\liam\AppData\Roaming\Typora\typora-user-images\image-20220210191444331.png)

To fix this go to the apache config file at `/etc/apache2/apache2.conf` and open it with `vi` or `nano`. Add the following lines to the end of the file.

```
ServerTokens Prod
ServerSignature Off
```

After doing this save and close the file and restart apache with `sudo systemctl restart apache2`. Now the version info won't show up in a packet capture anymore.

![image-20220210192801123](C:\Users\liam\AppData\Roaming\Typora\typora-user-images\image-20220210192801123.png)



### Disable Directory Browsing

Directory browning lets you view different directories stored within apaches `/var/www/html` directory. To test this we can add a new directory in `/var/www/html` by doing `sudo mkdir test` and then `sudo vi test.txt` and adding some words into that file. Now by going back to the apache web server page and type in `YOUR-SERVER-IP/test` it will then show the directory we added with the file. 

![image-20220210193728876](C:\Users\liam\AppData\Roaming\Typora\typora-user-images\image-20220210193728876.png)

To disable this open the config file in `/etc/apache2/apache2.conf` and find the lines below.

```
<Directory /var/www/>
        Options Indexes FollowSymLinks
        AllowOverride None
        Require all granted
</Directory>
```

Change the line:

`Options Indexes FollowSymLinks`

to:

`Options -Indexes +FollowSymLinks`

Save and close the file and then restart the Apache webserver.

`sudo systemctl restart apache2`

When trying to access the same directory we will now get a Forbidden message.

![image-20220210194510106](C:\Users\liam\AppData\Roaming\Typora\typora-user-images\image-20220210194510106.png)

### Etag

This allows remote attackers to obtain sensitive information like inode number, multipart MIME boundary, and child process through the Etag header. To prevent this vulnerability add the line below to the bottom of the apache config file:

`FileETag None`

Restart apache

`sudo systemctl restart apache2`

### HTTP Request Methods



## Check With Nikto

