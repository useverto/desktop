echo "
[Desktop Entry]
Terminal=false
Exec=/usr/bin/desktop
Icon=/usr/bin/verto_desktop.png
Type=Application
Categories=Graphics;
StartupNotify=true
Name=Verto Desktop
GenericName=Verto
" >> /usr/share/applications/verto_desktop.desktop

echo "[VERTO] Running first-time bootstrap"
cd /usr/bin/
./desktop