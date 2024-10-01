<script lang="ts">
    import Top from "$lib/components/Top.svelte";
    import Footer from "$lib/components/Footer.svelte";
    import OptionsList from "$lib/components/optionsList.svelte";
    import { onMount } from 'svelte';
    
    onMount(() => {
        const terminalElement = document.getElementById('terminal');
        const term = new Terminal();

        if (terminalElement) {
            term.open(terminalElement);
            term.write('$ ');

            term.onKey((key) => {
                const char = key.domEvent.key;
                if (char === 'Backspace') {
                    if (term.buffer.active.cursorX > 2) 
                        term.write('\b \b');
                    else if(term.buffer.active.cursorY > 0)
                        if(term.buffer.active.cursorX == 0) 
                            term.write('\b\b\b\b\b$ \b\b\b\b\b');
                        else 
                            term.write('\b \b');
                } else if (char === 'Enter') {
                    term.write('\r\n$ ');  
                } else if (char.length === 1) {
                    term.write(char);
                }
            });
        }
    });
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
        padding: 3px 0px 3px 10px;
        background-color: black;
    }

    #terminal::selection{
        background-color: white;
        color: black;
    }
</style>

<head>
    <link rel="stylesheet" href="../../../node_modules/@xterm/xterm/css/xterm.css" />
    <script src="../../../node_modules/@xterm/xterm/lib/xterm.js"></script>
</head>
<section>
    <Top showServers={true} isLoggedIn={true}/>
    <div>
        <nav>
            <OptionsList name={"name"} version={"version"} author={"author"}/>
        </nav>
        <div id="consoleContainer">
            <div id="terminal"></div>
        </div>
    </div>
    <Footer/>
</section>