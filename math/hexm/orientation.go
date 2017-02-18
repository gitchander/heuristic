package hexm

type Orientation int

//    Angled Orientation:
//             _
//          _-- --_
//       _--       --_
//    _--             --_
//   |                   |
//   |                   |
//   |                   |
//   |                   |
//   |                   |
//   |_                 _|
//     --_           _--
//        --_     _--
//           --_--
//

//      Flat Orientation:
//        ____________
//       /            \
//      /              \
//     /                \
//    /                  \
//   /                    \
//   \                    /
//    \                  /
//     \                /
//      \              /
//       \____________/
//

const (
	Angled Orientation = iota
	Flat
)
