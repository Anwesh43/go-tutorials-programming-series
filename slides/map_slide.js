var mainApp = angular.module('mainApp',[])
// mainApp.service('itemAnimator',()=>{
//     this.animateItems = function(items,scope) {
//
//     }
// })
mainApp.controller('mapController',function($scope){
    var items = ["What is Variadic function?","","how to use in golang?","A simple example"]
    items = items.map(function(item){
        var newItem =  {data:item,h:(window.innerWidth*8)/10,speed:window.innerWidth/20,t:0,limit:12}
        return newItem
    })
    $scope.items = []
    window.onkeydown = function(event) {
        if(event.keyCode == 32) {
            console.log(items)
            if(items.length>0) {
                var animatedItem = items.splice(0,1)[0]
                $scope.items.push(animatedItem)
                $scope.$apply()
                var interval = setInterval(function(){
                    animatedItem.h -= animatedItem.speed
                    animatedItem.t++
                    $scope.items[$scope.items.length-1] = animatedItem
                    $scope.$apply()
                    if(animatedItem.t >= animatedItem.limit) {
                        clearInterval(interval)
                    }
                },50)
            }
        }
    }
    //itemAnimator.animateItems(items,$scope)
})
