package field

/*
Field Type	Supported Interface
generic	IsNull/IsNotNull/Count/Eq/Neq/Gt/Gte/Lt/Lte/Like/Value/Sum/IfNull
int	Eq/Neq/Gt/Gte/Lt/Lte/In/NotIn/Between/NotBetween/Like/NotLike/Add/Sub/Mul/Div/Mod/FloorDiv/RightShift/LeftShift/BitXor/BitAnd/BitOr/BitFlip/Value/Zero/Sum/IfNull
uint	same with int
float	Eq/Neq/Gt/Gte/Lt/Lte/In/NotIn/Between/NotBetween/Like/NotLike/Add/Sub/Mul/Div/FloorDiv/Floor/Value/Zero/Sum/IfNull
string	Eq/Neq/Gt/Gte/Lt/Lte/Between/NotBetween/In/NotIn/Like/NotLike/Regexp/NotRegxp/FindInSet/FindInSetWith/Value/Zero/IfNull
bool	Not/Is/And/Or/Xor/BitXor/BitAnd/BitOr/Value/Zero
time	Eq/Neq/Gt/Gte/Lt/Lte/Between/NotBetween/In/NotIn/Add/Sub/Date/DateDiff/DateFormat/Now/CurDate/CurTime/DayName/MonthName/Month/Day/Hour/Minute/Second/MicroSecond/DayOfWeek/DayOfMonth/FromDays/FromUnixtime/Value/Zero/Sum/IfNull
以下是一些用法示例：
*/
