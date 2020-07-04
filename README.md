# golang-sciter-environment-setup-windows
This repository contains instructions to setup development environment for Sciter and golang on windows Machine.

Steps to setup environment for GOLANG with SCITER on Windows
============================================================


1. Download GoLang:  https://golang.org/dl/

	* Download the latest golang msi build, currently it is: https://golang.org/dl/go1.14.4.windows-amd64.msi
	
	* Make sure the golang directory is in your PATH 
		* Go to windows search & type "Environment Variables" -> Select 'Path' -> click 'Edit' -> Add your directory path (C:\Go\bin)
		* I installed it in 'c:\Go\bin', Ignore if it is already present
	
	* Test by running this command in cmd: go version
	
	

2. Download Git CLI: https://git-scm.com/download/
	
	* Add git to your PATH
		* Go to windows search & type "Environment Variables" -> Select 'Path' -> click 'Edit' -> Add your git directory path(C:\Program Files\Git\cmd)
		* I installed it in 'C:\Program Files\Git\cmd', Ignore if it is already present
	
	* Test it by running this command in cmd:  git --version

3. Download GCC compiler: https://jmeubank.github.io/tdm-gcc/
	
	* I installed the "MinGW-w64 based" msi packaage as I am running 64-bit windows machine.
	
	* Add it to your PATH
		* Go to windows search & type "Environment Variables" -> Select 'Path' -> click 'Edit' -> Add your GCC directory path(C:\Program Files\Git\cmd)
		* I installed it in 'C:\TDM-GCC-64\bin', Ignore if it is already present

4. Download sciter-SDK and add the sciter binary to your project folder according to your platform:
	
	* Official Website:  https://sciter.com/download/
	
	* extract the zip and copy sciter-sdk.zip\bin.win\x64\sciter.dll  to your project directory


5. Open your Code editor and Download dependencies for your project before building it with golang:
	
	* command :  ``go get package_name`` 

	* Basic dependencies for sciter:
		* ``go get github.com/sciter-sdk/go-sciter``
		* ``go get github.com/sciter-sdk/go-sciter/window``
	
	* Windows API wrapper :
 		* ``go get github.com/lxn/win``


