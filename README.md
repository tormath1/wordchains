## log book of the kata 19

1st day

First, I was like OMG it looks like to be the final boss of the maths problem. I REALLY love maths problem, I could send you a math problem I really like if you want.

So, I read like a dozen times kata19. Trying to find hints into the text. 

2nd day; I remember binary search tree, I have seen this 2 years ago. It was pretty nice to solve classification problem. But now, how could I use a binary tree ? I mean I do not know how to order the tree. But I think I have something.
*A few hours later* I should have many leafs to my tree... But the words need to have the same size. So first, I have preprocessed the data from the text file: remove bigger / shorter words. I am using pointer in every function, I do not want to copy an array each time I need to do something on it.

day 3: I've spent my evening watching mooc from MIT about search algorithm, and I've found something about breadth first search: it was exaclty what i'm looking for yesterday: a root with many leafs. And the guy was like : if you have a rubiks cube, you can find the shortest number of moves in order to solve it. Well, this is exactly our problem.
So now, I've made all the preliminiary work: preprocess data, ensure that words are in the dataset. I can implement this algorithm. and yes, finally it wrks ! i spent a few times trying to figure out how to store the previous word in order to build the path, I've finally decide to go with a list of a `node` struct.
well, I'am under the second for now. Let's see if we can go faster with a recursive way

a few minutes later: something I've learn on my experiences: early optimization is the gate of the hells. It works really like this, so we do not need to improve the efficience

ToDo: 

- [ ] add an interface in order to implement many algorithms to solve the problem
- [ ] add a CLI for the user
