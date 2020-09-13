# creates the following directory structure
# verto_desktop.app
# └── Contents
#     ├── Info.plist
#     ├── MacOS
#     |    └── verto_desktop
#     └── Resources
#         └── verto_desktop.icns

mkdir verto_desktop.app
mkdir verto_desktop.app/Contents
mkdir verto_desktop.app/Contents/MacOS
mkdir verto_desktop.app/Contents/Resources
cp desktop verto_desktop.app/Contents/MacOS/verto_desktop
cp Info.plist verto_desktop.app/
cp assets/verto_desktop.icns verto_desktop.app/Contents/Resources/
