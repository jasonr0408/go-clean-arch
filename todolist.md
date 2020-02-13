1. 建立db，匯入資料 ok
2. 建立玩家假資料 ok

3. getGameLink實作 
   1. input: gameId、user name、lang
   2. output:           gameUrl = sprintf("%s%s/%s/?game=%s&operator=%s&currency[0]=required&locale=%s&mode=%s&launchtoken=%s&agentId=%s",
              gameInfo['GameUrlH5'], gameInfo['UserName'],"s" + gameInfo['betlogcode'], gameInfo['gameName'], gameInfo['UserName'], url_lang, gameInfo['SettlementMode'],param.launchtoken, gameInfo['agent_relay_id']);
   
   組Launch Token跟player info
   回傳

4. 模擬Game server跟平台API call auth api
   比對launch Token
   回傳

5. 提供魚機跟Slot遊戲 api