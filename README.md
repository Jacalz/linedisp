# linedisp
Display lines in Fyne using linear algebra! The idea of this project is to show how linear algebra can be used for computer graphics in Fyne.

![3d-matrix-lines-v3](https://user-images.githubusercontent.com/25466657/143089168-dc190181-496f-4c5b-8077-0cdbb547ed88.gif)

## How it works
Multiple points are added together to form a coordinate matrix. We then tell Fyne to draw lines between all of the positions within the columns of the coordinate matrix. We can then simply use matrix transformations to modify the coordinate matrix and thus also the view on screen.

For rotation with the mouse, it uses two threedimensional [rotation matricies](https://en.wikipedia.org/wiki/Rotation_matrix#Basic_rotations), those for X and Y rotation, multiplied together to a single matrix. This allows the view to be roteted only by multiplying the coordinate matrix and telling Fyne to re-draw it.

For scaling with the scroll whell, it uses a [scaling matrix](https://en.wikipedia.org/wiki/Scaling_(geometry)#Matrix_representation) to make the view larger or smaller on the screen. Again, it is as simple as multiplying the the coordinate matrix and telling Fyne to re-draw it.

## Built using
This project is built using the following projects:
- Fyne: https://github.com/fyne-io/fyne
- Linalg: https://github.com/Jacalz/linalg

## License

Linedisp is licensed under the `BSD 3-Clause License` and freely avaliable to all of those that wish to use it.
