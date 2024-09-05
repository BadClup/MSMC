<script lang="ts">
    import Switch from "$lib/components/Switch.svelte";
    export let name = "Failed to load name";
    export let version = "Failed to load version information";
    export let author = "Failed to load author information";
</script>
<style>
    #optionsList{
        display: inline-block;
        box-sizing: border-box;
        width: 220px;
        height: 100%;
        padding: 0;
        background-color: rgb(129, 155, 182);
    }
    
    #statusMargin{
        height: 100%;
        width: 11px;
        background-color: navy;
        box-sizing: border-box;
        display: inline-block;
        margin: 0;
    }

    #statusContainer{
        width: 100%;
        height: 140px;
        background-color: rgb(59, 59, 207);
        display: inline-block;
        box-sizing: border-box;
        margin: 0;
        text-align: left;
    }

    #status{
        height: 90px;
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
        height: 34px;
        width: 60px;
        text-align: right;
        float: right;
        padding-top: 53px;
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
        padding-top: 25px;
    }

    p{
        margin: 0;
        padding: 0;
    }
</style>

<div id="optionsList">
    <div id="statusContainer">
        <div id="statusMargin"></div>
        <div id="status">
            <p>Name: {name}</p>
            <p>Version: {version}</p>
            <p>Author: {author}</p>
        </div> 
        <div id="switchWrapper"><Switch/></div>
    </div>
    <ul>
        <li><a href="/configuration">Server</a></li>
        <li><a href="/configuration/settings">Settings</a></li>
        <li><a href="/configuration/console">Console</a></li>
        <li><a href="/configuration/logs">Logs</a></li>
        <li><a href="/configuration/players">Players</a></li>
        <li><a href="/configuration/world">World</a></li>
        <li><a href="/configuration/backups">Backups</a></li>
        <li><a href="/configuration/access">Access</a></li>
    </ul>
</div>
