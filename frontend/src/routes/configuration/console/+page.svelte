<script lang="ts">
    import Top from "$lib/components/Top.svelte";
    import Footer from "$lib/components/Footer.svelte";
    import OptionsList from "$lib/components/optionsList.svelte";
</script>
<style>
    section{
        display: flex;
        flex-direction: column;
        height: 100vh;
        box-sizing: border-box;
        padding: 0;
        background-color: rgb(88, 88, 88);
        color: white;
    }

    section > div{
        width: 100%;
        flex: 1;
    }

    nav{
        display: inline-block;
        float: left;
        height: 100%;
        width: 240px;
    }

    #consoleContainer{
        display: inline-block;
        box-sizing: border-box;
        height: 100%;
        width: calc(100% - 240px);
        float: right;
        flex: 1;
        padding: 3%;
    }

    #terminal{
        width: 100%;
        height: 100%;
    }
</style>

<section>
    <link rel="stylesheet" href="../../../node_modules/@xterm/xterm/css/xterm.css"/>
    <script src="../../../node_modules/@xterm/xterm/lib/xterm.js"></script>

    <Top showServers={true} isLoggedIn={true}/>
    <div>
        <nav>
            <OptionsList name={"name"} version={"version"} author={"author"}/>
        </nav>
        <div id="consoleContainer">
            <div id="terminal"></div>
            <script>
                var term = new Terminal();
                term.open(document.getElementById('terminal'));
                term.write('MSMC console $ ');
            </script>
        </div>
    </div>
    <Footer/>
</section>