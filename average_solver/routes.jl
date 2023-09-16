using Genie.Router, Genie.Exceptions, Genie.Renderer.Html, Genie.Requests
using AverageSolver
using JSON3
using StructTypes
using UUIDs


struct Item
    id::Int
    size::Int
    value::Int
end

struct Knapsack
    bagSize::Int
    items::Vector{Item}
end

StructTypes.StructType(::Type{Knapsack}) = StructTypes.Struct()


route("/average-solver/solve", method = POST) do
    uuidReq = string(UUIDs.uuid4())

    knapsack = JSON3.read(rawpayload(), Knapsack)

    @info  "New request id: $uuidReq with $(knapsack.bagSize) bagsize and $(length(knapsack.items)) items."

    start_time = now()
    max_value, selected_items = AverageSolver.solve(knapsack.bagSize, knapsack.items)
    duration = (now() - start_time) / Dates.Millisecond(1000)

    @info "Solved request id: $uuidReq with max_value=$max_value and items=$(JSON3.write(selected_items)). Took $duration s"

    # Build the JSON data
    json_data = Dict(
        "max_value" => max_value,
        "items" => selected_items,
        "duration" => duration
    )

    JSON3.write(json_data)
end

route("/average-solver/health") do
  html("average solver is healthy")
end
