(pwd() != @__DIR__) && cd(@__DIR__) # allow starting app from bin/ dir

using AverageSolver
const UserApp = AverageSolver
AverageSolver.main()
