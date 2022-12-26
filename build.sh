#!/bin/sh

test -f Spirit-Installer.dmg && rm Spirit-Installer.dmg

wails build -u -platform darwin/universal
create-dmg \
    --volname "Spirit Installer" \
    --volicon "build/bin/Spirit.app/Contents/Resources/iconfile.icns" \
    --icon-size 75 \
    --window-size 600 400 \
    --icon "Spirit.app" 200 170 \
    --app-drop-link 400 170 \
    --hide-extension "Spirit.app" \
    "Spirit-Installer.dmg" \
    "build/bin/"
