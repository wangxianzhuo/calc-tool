# calc-tool

> calculate tools

## Tools

### Modbus tool

- modbus-rtu crc16 check
- modbus-rtu instruction check
- modbus-rtu instruction parse
- modbus-rtu instruction generate with check code
- modbus-rtu response check

### Hydropower

- ~cubic spline~
- curve sluice open degree
- straight sluice open degree
- sluice overflow
- dam overflow
- hydro-unit used flow
- water head
- water head los

### Numberic/Circle

- get z by x & y
- get z from io.Reader

### 机组特性曲线展示

曲线数据由多行组成，每行数据由四列构成。这四列分别是：

1 | z | x | y
--|---|---|---

例如：

1 | z | x | y
--|:-:|:-:|:-:
1|10|-0.9999999999999999|-1.4901161193847656e-08
1|10|-0.9999999999999999|1.4901161193847656e-08
1|9|-0.8999999999999999|0.43588989435406755
1|8|-0.7999999999999999|0.6000000000000001
