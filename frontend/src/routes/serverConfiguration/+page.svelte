<script lang="ts">
    import Top from "./../Top.svelte";
    import Footer from "./../Footer.svelte";
    import Switch from "./../Switch.svelte";
</script>
<style>
    #container{
        display: flex;
        flex-direction: column;
        height: 100vh;
        box-sizing: border-box;
        padding: 0;
        background-color: rgb(88, 88, 88);
    }

    #options{
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
        display: inline-block;
        box-sizing: border-box;
        text-align: left;
        position: relative;
        bottom: 80px;
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
</style>

<div id="container">
    <Top logout={true}/>
    <div id="options">
        <div id="statusContainer">
            <div id="statusMargin"></div>
            <div id="status">
                Nazwa<br>
                Wersja<br>
                Autor
            </div>
            <Switch/>
        </div>
        <ul>
            <li>Settings</li><br>
            <li>Console</li><br>
            <li>Logs</li><br>
            <li>Players</li><br>
            <li>World</li><br>
            <li>Backups</li><br>
            <li>Access</li><br>
        </ul>
    </div>
    <Footer/>
</div>