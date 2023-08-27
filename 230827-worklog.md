# 2023-08-27    16:03
=====================

    $ cd ubuntu18
    $ vagrant init ubuntu/bionic64
----------------------------------
    $ vim Vagrantfile
# vi: set ft=ruby :
Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/bionic64"

  # Create a private network, which allows host-only access to the machine
  # using a specific IP.
  # config.vm.network "private_network", ip: "192.168.33.10"

   config.vm.network "public_network"
   config.vm.synced_folder "../shared_directory", "/shared_directory"
   config.vm.synced_folder ".", "/vagrant", disabled: true
   config.vm.provider "virtualbox" do |vb|
      vb.memory = "1024"
      vb.cpus = 1
   end

   config.vm.provision "shell", inline: <<-SHELL
     apt-get update
     apt-get install -y mc tree zip unzip
     cd ~
     wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
     rm -rf /usr/local/go && tar -C /usr/local -xzvf go1.21.0.linux-amd64.tar.gz
     echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile
     go version
   SHELL
end
----------------------------------


# требуется создать ярлык для SSH ключа
    default: Key inserted! Disconnecting and reconnecting using new SSH key...
The private key to connect to the machine via SSH must be owned
by the user running Vagrant. This is a strict requirement from
SSH itself. Please fix the following key to be owned by the user
running Vagrant:

/mnt/.../ubuntu18/.vagrant/machines/default/virtualbox/private_key


    $ ln -sr vagrant_ubuntu18_private_key /mnt/.../ubuntu18/.vagrant/machines/default/virtualbox/private_key
    $ vagrant reload --provision

# 2023-0827 17:53
=================
    ubuntu18$ vagrant ssh
    $ cd ../shared_directory/scratch
    $ go run main.go run echo hello Dnipro
Hello, Azrubael!
running [echo hello Dnipro] as PID 37626
    $ go run main.go run /bin/bash
Hello, Azrubael!
running [/bin/bash] as PID 37556
    $ ps
    PID TTY          TIME CMD
   2605 pts/0    00:00:00 bash
  37632 pts/0    00:00:00 ps
