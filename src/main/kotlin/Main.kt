package org.example

import org.example.RunSmoke.runCommand
import java.util.concurrent.Executors
import java.util.concurrent.TimeUnit

fun main() {
    val executor = Executors.newSingleThreadScheduledExecutor()

    val task = Runnable { runCommand("sudo -n true") }
    // runCommand("flatpak run com.google.Chrome") }
    // runCommand("sudo systemctl start smokeping")

    executor.scheduleAtFixedRate(task, 0, 1, TimeUnit.MINUTES)

}
