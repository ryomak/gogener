workflow "Golang workflow" {
  on = "push"
  resolves = ["Test"]
}
 
action "GolangCI-Lint" {
  uses = "./.github/actions/go"
  args = "lint"
}
 
action "Test" {
  needs = ["GolangCI-Lint"]
  uses = "./.github/actions/go"
  args = "test"
}