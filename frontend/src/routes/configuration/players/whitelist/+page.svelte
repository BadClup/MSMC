<script lang="ts">
    import Top from "$lib/components/Top.svelte";
    import Footer from "$lib/components/Footer.svelte";
    import OptionsList from "$lib/components/optionsList.svelte";
    import remove from '$lib/images/remove.svg';

    interface Player{
        name: string;
    }

    const whitelist: Player[] = [{name: "player1"}, {name: "player2"}, {name: "player3"}, {name: "player4"}, {name: "player5"}]
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

    main{
        display: inline-block;
        box-sizing: border-box;
        width: calc(100% - 240px);
        height: 100%;
        padding: 2% 3%;
    }

    h1{
        position: relative;
        text-align: center;
        font-weight: bold;
        font-size: 50px;
        color: inherit;
        margin: 0;
        margin-bottom: 40px;
    }

    .playerWrapper{
        box-sizing: border-box;
        display: inline-block;
        width: 100%;
        height: 50px;
        margin-top: 18px;
        background-color: rgb(48, 48, 48);
    }

    .playerInfo{
        display: inline-block;
        background-color: rgb(48, 48, 48);
        color: white;
        font-weight: 600;
        font-size: 20px;
        padding: 13px;
    }

    .playerInfo::before{
        content: url($lib/images/steve.png);
        margin-right: 10px;
    }

    .buttonsContainer{
        display: inline-block;
        float: right;
        background-color: rgb(48, 48, 48);
        color: white;
        padding: 10px 15px 0px;
    }

    button{
        width: 30px;
        height: 30px;
        background-color: transparent;
        border: 0;
        color: white;
    }
</style>

<section>
    <Top showServers={true} isLoggedIn={true}/>
    <div>
        <nav><OptionsList name={"name"} version={"version"} author={"author"}/></nav>
        <main>
            <h1>Whitelist</h1>
            {#each whitelist as player}
            <div class="playerWrapper">
                <div class="playerInfo">{player.name}</div>
                <div class="buttonsContainer">
                    <button><img src={remove} alt="Remove"></button>
                </div>
            </div>
        {/each}
        </main>
    </div>
    <Footer/>
</section>