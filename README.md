# AWS Knapsack Microservices Project
# Project Overview & Description
The goal of this project is to create a Proof Of Concept (PoC) using different AWS services for a Knapsack greedy problem implementation.
This implementation will follow an Microservices Event driven architecture to test simultaneously different types of algorithms implemented in different ways, languages and using different (AWS) technologies.
For more information on the architecture, please refer to [the architecture page](architecture.md) that list all the different iteration and evolution of the architectural and solution design.

Finally, this project should be able to be used and reproduced following an 'Open/Closed Principle' ([OCP](https://stackify.com/solid-design-open-closed-principle/)) with ability to easily add more algorithms and take advantage of logs and analytics data generated.
It will also use modern CI/CD technologies (see the full list [below](#List-of-Technologies-&-services-used)) to automatically deploy, run and monitor each services/modules of the whole solution.

# List of Technologies & services used
* Go (for problem generator)
* GitHub Actions (for CI/CD)
* TBD

# RoadMap
For a complete roadmap please refer to the [Project Tab](https://github.com/users/PERES-Richard/projects/1/views/1) of this project.

- [x] Create a simple generator of knapsack problems (list of items + bag size)  
- [ ] Create a first simple MVP of a naive solver  
- [ ] Create a MVP UI to launch a complete flow (generate a new problem, send it to solver and print result)  
