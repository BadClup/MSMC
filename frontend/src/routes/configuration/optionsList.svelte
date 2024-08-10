<script lang="ts">
    import Switch from "$lib/components/Switch.svelte";
    export let name = "Failed to load name";
    export let version = "Failed to load version information";
    export let author = "Failed to load author information";
</script>
<style>
    section{
        display: inline-block;
        box-sizing: border-box;
        width: 240px;
        height: 100%;
        padding: 0;
        background-color: rgb(129, 155, 182);
    }

    section > div{
        width: 100%;
        height: 140px;
        background-color: rgb(59, 59, 207);
        display: inline-block;
        box-sizing: border-box;
        margin: 0;
        text-align: left;
    }

    #statusMargin{
        height: 100%;
        width: 11px;
        background-color: navy;
        box-sizing: border-box;
        display: inline-block;
        margin: 0;
    }

    #status{
        height: 100%;
        display: inline-block;
        box-sizing: border-box;
        text-align: left;
        position: relative;
        bottom: 50px;
        margin: 0;
        padding-left: 4px;
        font-size: 18px;
    }

    #switchWrapper{
        display: block;
        height: 100%;
        width: 60px;
        text-align: right;
        float: right;
        padding-top: 53px;
        padding-right: 6px;
    }

    ul{
        margin: 0;
        padding: 0;
        border-bottom: solid 1px black;
    }

    li{
        display: inline-block;
        list-style-type: none; 
        padding: 0; 
        margin: 0; 
        box-sizing: border-box;
        text-align: center;
        width: 100%;
        height: 80px;
        background-color: rgb(129, 155, 182);
        font-size: 24px;
        font-weight: bold;
        color: white;
        border: solid 1px black;
        padding-top: 25px;
    }

    li:hover{
        background-color: rgb(94, 159, 228);
    }

    a{
        display: inline-block;
        width: 100%;
        height: 100%;
        text-decoration: none;
        color: inherit;
    }
</style>

<section>
    <div>
        <div id="statusMargin"></div>
        <div id="status">
            Name: {name}<br>
            Version: {version}<br>
            Author: {author}
        </div> 
        <div id="switchWrapper"><Switch/></div>
    </div>
    <ul>
        <li><a href="/configuration">Server</a></li><br>
        <li><a href="/configuration/settings">Settings</a></li><br>
        <li><a href="/configuration/console">Console</a></li><br>
        <li><a href="/configuration/logs">Logs</a></li><br>
        <li><a href="/configuration/players">Players</a></li><br>
        <li><a href="/configuration/world">World</a></li><br>
        <li><a href="/configuration/backups">Backups</a></li><br>
        <li><a href="/configuration/access">Access</a></li><br>
    </ul>
</section>
