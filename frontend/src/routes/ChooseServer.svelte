<script lang="js">
    /*funtion createServerObject(serverName, serverVersion, serverStatus){
        this.name = serverName;
        this.version = serverVersion;
        this.status = serverStatus;
    }

    let Servers = [new createServerObject("serwer1", "1.19.2", 0), new createServerObject("serwer2", "1.16.4", 1), new createServerObject("serwer3", "1.20.1", 0)]*/
    let Servers = [["server1", "1.19.2", 0], ["server2", "1.16.4", 1], ["server3", "1.20.1", 0]]
</script>

<style>
    #ChooseServer{
        text-align: Center;
        font-size: 2em;
        background-color: rgb(88, 88, 88);
        color: rgb(240,240,240);
        flex: 1;
    }

    .ServerOption{
        display: inline-block;
        font-size: 1em;
        height: 100px;
        width: 200px;
    }
</style>

<div id="ChooseServer">
    Your Servers<br><br>
    {#each Servers as Server}
        <div class="ServerOption">
            {Server[0]}<br>{Server[1]}
        </div><br><br>
    {/each}
</div>