## kata-19: words chain

Try to link two words by a chain of words with one different character.

Example: `code` -> `ruby`: code rode robe rube ruby

## log book of the exercice

### 1st day

First, I was like OMG it looks like to be the final boss of the maths problem. I REALLY love maths problem, I could send you a math problem that I really like if you want. I read like a dozen times the kata 19: trying to find hints into the text. 

### 2nd day

I am figure out that I need to use a graph or something like that.I remember binary search tree, I have seen this 2 years ago. It was pretty nice in order to solve classification problem. Now, how could I use a binary tree ? I mean I do not know how to order the tree. But I think I have something.
*A few hours later* I should have many leafs to my tree... And the words need to have the same size. First, I need to prepare the data from the text file: remove words with a different size from the two given inputs.

### 3rd day

I have spent my evening watching MOOC from MIT about search algorithm, and I have found something about Breadth First Search (BFS): it was exaclty what I am looking for yesterday: a root with many leafs. The lecturer was like : if you have a rubiks cube, you can find the shortest number of moves in order to solve it: this is exactly our problem.
I have made all the requirements work: preprocess data, ensure that inputs are in the dataset. I can implement this algorithm. Yes, finally it works ! I spent a few times trying to figure out how to store the previous word in order to build the path, I have finally decide to use a list of `node`.
I am under the second for now. Let's see if we can go faster with a recursive way.

*a few minutes later* Something I have learnt on my experiences: early optimization is the gate of the hells. It works really well like this, so we do not need to improve the efficience.

## ToDo: 

- [ ] add an interface in order to implement many algorithms to solve the problem
- [ ] add a CLI for the user
