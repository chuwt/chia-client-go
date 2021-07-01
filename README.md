# chia-client
go编写的chia rpc client，主要包含`余额查询`和`交易发送`

## rpc
1. 获取coin列表
- 直接调用节点的/get_coin_records_by_puzzle_hash 方法
   
2. 发送签名的交易
- 直接调用节点的/push_tx 方法
   
3. 签名未签名的交易
- 首先通过[chia-tx](https://github.com/chuwt/chia-tx)生成未签名的
交易，然后用此方法将交易签名，签名后的交易可直接通过`signTx`发送出去