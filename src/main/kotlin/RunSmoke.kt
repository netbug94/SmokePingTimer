package org.example

object RunSmoke {
    fun runCommand(command: String) {
        val process = ProcessBuilder(*command.split(" ").toTypedArray())
            .redirectOutput(ProcessBuilder.Redirect.INHERIT)
            .redirectError(ProcessBuilder.Redirect.INHERIT)
            .start()

        process.waitFor()
    }
}
