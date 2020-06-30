var cron = require('node-cron');

//execute every 7 min
cron.schedule('*/10 * * * *', function(){
    var shell = require('./child_helper');
    var commandList = [
        "go run main.go"
    ]
    shell.series(commandList , function(err){
        console.log('done')
    });
});
