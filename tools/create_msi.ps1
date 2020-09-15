# compile each .wxs into a .wixobj object using candle.exe
exec { &"candle.exe" -out "Verto.wixobj" "Verto.wxs" } 
# link all the .wixobj objects into an MSI installer using light.exe
exec { &"light.exe" -b "./desktop.exe" "Verto.wixobj" -out "Verto.msi" }
