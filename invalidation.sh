#!/bin/bash

# Makes ArchiveInvalidation.txt

FONV_DIR="${HOME}/.steam/steam/steamapps/common/Fallout New Vegas"
TEXS_DIR="${FONV_DIR}/Data/Textures"

find "${TEXS_DIR}" -type f | cut -c "$((${#TEXS_DIR}+2))-" > "${FONV_DIR}/ArchiveInvalidation.txt"
