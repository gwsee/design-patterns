var Singleton = (function() {
    // 私有的单例实例引用
    var instance;

    // 定义一个初始化函数，它将作为真正的构造函数
    function init() {
        // 单例对象的私有方法和属性
        function privateMethod() {
            console.log("I am private");
        }

        var privateVariable = "I'm also private";

        var privateRandomNumber = Math.random();

        return {
            // 公共方法和属性
            publicMethod: function() {
                console.log("The public can see me!");
            },
            publicProperty: "I am also public",
            getRandomNumber: function() {
                return privateRandomNumber;
            }
        };
    }

    return {
        // 如果存在实例则返回该实例，如果不存在则创建新实例
        getInstance: function() {
            if (!instance) {
                // 创建 Singleton 实例
                instance = init();
            }
            return instance;
        }
    };
})();

var singleA = Singleton.getInstance();
var singleB = Singleton.getInstance();

console.log(singleA === singleB); // 输出：true
console.log(singleA.getRandomNumber() === singleB.getRandomNumber()); // 输出：true
