// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// Version: 1.3.0.0


class Order: Comparable {
//    var id: Int = 0
//    var symbol: String = ""
//    //var side: OrderSide = OrderSide()
//   / var type: OrderType = OrderType()
//    var price: Double = 0.0
//    var volume: Double = 0.0
//
//    init() { }
//    init(id: Int, symbol: String, side: OrderSide, type: OrderType, price: Double, volume: Double) {
//        self.id = id
//        self.symbol = symbol
//        self.side = side
//        self.type = type
//        self.price = price
//        self.volume = volume
//    }
//
//    init(other: Order) {
//        self.id = other.id
//        self.symbol = other.symbol
//        self.side = other.side
//        self.type = other.type
//        self.price = other.price
//        self.volume = other.volume
//    }
//
//    func clone() throws -> Order {
//        // Serialize the struct to the FBE stream
//        let writer = OrderModel()
//        try _ = writer.serialize(value: self)
//
//        // Deserialize the struct from the FBE stream
//        let reader = OrderModel()
//        reader.attach(buffer: writer.buffer)
//        return reader.deserialize()
//    }
//
//    static func < (lhs: Order, rhs: Order) -> Bool {
//        if !(lhs.id < rhs.id) { return false }
//
//        return true
//    }
//
//    static func == (lhs: Order, rhs: Order) -> Bool {
//        if !(lhs.id < rhs.id) { return false }
//
//        return true
//    }
//
//    func hashCode() -> Int {
//        var hash = 17
//        hash = hash * 31 + id.hashValue
//        return hash
//    }
//
//    func toString() -> String {
//        var sb = String()
//        sb.append("Order(")
//        sb.append("id="); sb.append(id.description)
//        sb.append(",symbol="); sb.append("\""); sb.append(symbol); sb.append("\"")
//        sb.append(",side="); sb.append(side.description)
//        sb.append(",type="); sb.append(type.description)
//        sb.append(",price="); sb.append(price.description)
//        sb.append(",volume="); sb.append(volume.description)
//        sb.append(")")
//        return sb
//    }
//}
