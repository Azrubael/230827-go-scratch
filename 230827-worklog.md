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
     cd Downloads
     apt-get update
     apt install -y mc tree zip unzip
     wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
     rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
     apt install -y go1.21.0.linux-amd64.tar.gz
     export PATH=$PATH:/usr/local/go/bin
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

/mnt/SSDATA/CODE/DevOpsCompl20/230824-VagrantVMs/ubuntu18/.vagrant/machines/default/virtualbox/private_key


    $ ln -sr $HOME/.ssh/vagrant_ubuntu18_private_key /mnt/SSDATA/CODE/DevOpsCompl20/230824-VagrantVMs/ubuntu18/.vagrant/machines/default/virtualbox/private_key
    $ vagrant reload --provision

