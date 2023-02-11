# linedisp
Using coordinate matrices and linear algebra, we can project a 3D model on screen. A linear map using rotation and scale matrices allows support for both rotation and zoom. The lines are drawn entirely using the regular `canvas.Line` objects in Fyne. 

[matrix-display.webm](https://user-images.githubusercontent.com/25466657/218252020-63855a3f-8340-4666-b008-9864070a1143.webm)

## How it works
Multiple points are added together to form a coordinate matrix. We then tell Fyne to draw lines between all of the positions within the columns of the coordinate matrix. We can then simply use matrix transformations to modify the coordinate matrix and thus also the view on screen.

For rotation with the mouse, it uses two threedimensional [rotation matricies](https://en.wikipedia.org/wiki/Rotation_matrix#Basic_rotations), those for X and Y rotation, multiplied together to a single matrix. This allows the view to be roteted only by multiplying one coordinate matrix (instead of two times) and then telling Fyne to re-draw it.

For scaling with the scroll whell, it uses a [scaling matrix](https://en.wikipedia.org/wiki/Scaling_(geometry)#Matrix_representation) to make the view larger or smaller on the screen. Again, it is as simple as multiplying the the coordinate matrix and telling Fyne to re-draw it.

## Built using
This project is built using the following projects:
- Fyne: https://github.com/fyne-io/fyne
- Linalg: https://github.com/Jacalz/linalg

## License

Linedisp is licensed under the `BSD 3-Clause License` and freely avaliable to all of those that wish to use it.
