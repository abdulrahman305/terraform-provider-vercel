# Environment variables can be identified by their ID, or by their key and target.
# The ID can be taken from the network tab inside developer tools, while you are on the project page.
data "vercel_shared_environment_variable" "example" {
  id = "xxxxxxxxxxxxxxx"
}

# Alternatively, you can use the key and target to identify the environment variable.
# Note that all `target`s must be specified for a match to be found.
data "vercel_shared_environment_variable" "example_by_key_and_target" {
  key    = "MY_ENV_VAR"
  target = ["production", "preview"]
}
