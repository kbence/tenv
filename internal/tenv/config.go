package tenv

import "os"

const FileMode = 0755
const DirectoryMode = 0755

var TeleportEnvHomeDirectory = os.ExpandEnv("$HOME/.tenv")
var TeleportEnvVersionDirectory = TeleportEnvHomeDirectory + "/versions"

var TeleportBinaryNames = []string{"teleport", "tsh", "tctl", "tbot"}
