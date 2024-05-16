class Handler {
    constructor(name, condition) {
      this.name = name;
      this.condition = condition; // 新增条件属性来决定是否处理请求
      this.nextHandler = null;
    }
  
    setNext(handler) {
      this.nextHandler = handler;
    }
  
    handleRequest(request) {
      // 检查是否满足条件来处理请求
      if (this.condition === request) {
        console.log(`${this.name} handles request: ${request}`);
      } else if (this.nextHandler) {
        console.log(`${this.name} passes request to ${this.nextHandler.name}`);
        this.nextHandler.handleRequest(request);
      } else {
        console.log("End of chain, request not handled");
      }
    }
  }
  
  const handlerA = new Handler('Handler A', 'request for A');
  const handlerB = new Handler('Handler B', 'request for B');
  const handlerC = new Handler('Handler C', 'request for C');
  
  handlerA.setNext(handlerB);
  handlerB.setNext(handlerC);
  
  // 发送不同的请求
  handlerA.handleRequest('request for A');
  // 输出: Handler A handles request: request for A
  
  handlerA.handleRequest('request for B');
  // 输出: Handler A passes request to Handler B
  //       Handler B handles request: request for B
  
  handlerA.handleRequest('request for C');
  // 输出: Handler A passes request to Handler B
  //       Handler B passes request to Handler C
  //       Handler C handles request: request for C
  
  handlerA.handleRequest('request for D');
  // 输出: Handler A passes request to Handler B
  //       Handler B passes request to Handler C
  //       End of chain, request not handled
  