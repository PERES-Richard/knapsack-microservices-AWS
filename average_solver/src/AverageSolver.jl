module AverageSolver

using Genie

const up = Genie.up
export up

function main()
  Genie.genie(; context = @__MODULE__)
end

function solve(bagSize, items)
    # Calculate the "averageValue" for each item and store it in a new array
    average_values = [item.value / item.size for item in items]

    # Create an array of tuples where each tuple contains the item index and its averageValue
    item_indices_and_values = [(i, avg) for (i, avg) in enumerate(average_values)]

    # Sort the items by averageValue in descending order
    sorted_items = sort(item_indices_and_values, by=x -> x[2], rev=true)

    selected_items = []
    total_size = 0

    for (item_index, _) in sorted_items
        item = items[item_index]

        # Check if adding the item exceeds the bagSize
        if total_size + item.size <= bagSize
            push!(selected_items, item)
            total_size += item.size
        else
            break  # Stop adding items if the bagSize is exceeded
        end
    end

    max_values = sum(item.value for item in selected_items)
    return max_values, selected_items
end


end
