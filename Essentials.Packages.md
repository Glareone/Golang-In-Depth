# Packages
![image](https://github.com/user-attachments/assets/1636ae2b-b000-47e7-bf15-83e38e3d2f2c)

* All code is organized in modules.
* you can have any number of modules you want
* you can import the methods and use them from other packages written by other developers or located in Go standard libraries using import (for instance, import "fmt")

# The "Main" Package
* main module is a special module
* it has the special entry point in your application

# Modules
![image](https://github.com/user-attachments/assets/f0e4c86c-39cd-4c82-b938-925b686393e6)
* Any module can consist of many packages. In C# world Module is equivalent to Project
* In order to build the app or use it along with "go run app.go" command - you need to initialized the Module

## Module Initialization
![image](https://github.com/user-attachments/assets/2d454a21-1ab4-4f36-94a2-aed59a58241e)

* "example.com/m" is a dummy name, usually you can place here the address where you want to distribude or consume your library
