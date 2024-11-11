# mandelbrot-toy
Mandelbrot drawing program written in go using fyne.io 


# Running on Windows
## msys-x64
% go build -v main.go
% go run -v main.go


# Building .apk for Android Studio
## msys-x64 
% fyne package -os android -appID com.example.mandelbrottoy -icon assets/mandelbrot-toy.png
% fyne package -os windows -icon assets/mandelbrot-toy.png
 
# Installing .apk to Android Studio
Drag Mandelbrot_Toy.apk to emulator icon

# Installaion on Ubuntu 
## - Assumptions
 * GOHOME = /home/craigwarner
%sudo apt install libgl1-mesa-dev xorg-dev
%go run -v main.go

# Intalling  
1) sudo /home/craigwarner/go/bin/fyne get github.com/craig-warner/mandelbrot-toy


