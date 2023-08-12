<script lang="ts">
  import type { CalendarEvent } from "../models/CalendarEvent";
  import { Filters, getTermsFromFilter } from "../models/Filters";
  import CalendarEventCard from "./CalendarEventCard.svelte";
  import Chip from "./Chip.svelte";

  export let events: CalendarEvent[];
  let filteredEvents: CalendarEvent[] = events;

  let searchTerm: string = "";

  $: filteredEvents = events.filter(
    (event) =>
      event.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      event.day.toLowerCase().includes(searchTerm.toLowerCase()) ||
      event.start_date.toLowerCase().includes(searchTerm.toLowerCase()) ||
      (event.end_date &&
        event.end_date.toLowerCase().includes(searchTerm.toLowerCase())) ||
      event.week.toLowerCase().includes(searchTerm.toLowerCase())
  );

  let selectedFilter: Filters | null = null;

  function getDayParts(day: string) {
    return day.split("-");
  }

  function getStartDay(day: string) {
    return getDayParts(day)[0];
  }

  function getEndDay(day: string) {
    return getDayParts(day).length > 1 ? getDayParts(day)[1] : "";
  }

  function onFilterClicked(filter: Filters) {
    if (filter == selectedFilter) {
      selectedFilter = null;
      filteredEvents = events;
    } else {
      selectedFilter = filter;
      let filterTerms = getTermsFromFilter(filter);

      filteredEvents = events.filter((event) =>
        filterTerms.some((term) => event.name.toLowerCase().includes(term))
      );
    }
  }
</script>

<div class="center">
  <input
    class="search-field"
    type="text"
    placeholder="Search for events"
    bind:value={searchTerm}
  />
</div>

<div class="custom-filters">
  {#each Object.entries(Filters) as [_, filter]}
    <Chip
      text={filter}
      isSelected={selectedFilter == filter}
      onClick={() => onFilterClicked(filter)}
    />
  {/each}
</div>

{#if filteredEvents.length > 0}
  <ul class="link-card-grid">
    {#each filteredEvents as event}
      <CalendarEventCard
        title={event.name}
        body={`${getStartDay(event.day)} ${event.start_date} ${
          getEndDay(event.day) == "" ? "" : "-"
        } ${getEndDay(event.day)} ${
          event.end_date != null ? event.end_date : ""
        }`}
        isHighlighted={false}
      />
    {/each}
  </ul>
{:else}
  <div class="center">
    <p class="not-found-text">No events found 😕</p>
  </div>
{/if}

<style>
  .link-card-grid {
    display: grid;
    grid-template-columns: repeat(auto, minmax(24ch, 1fr));
    gap: 1rem;
    padding: 0;
  }

  .center {
    display: flex;
    justify-content: center;
  }

  .search-field {
    width: 100%;
    padding: 0.5rem;
    border: 1px solid var(--color-border);
    border-radius: 0.5rem;

    font-size: 1rem;
  }

  .not-found-text {
    font-size: 1.5rem;
    font-weight: 500;
  }

  .custom-filters {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    gap: 1rem;
    margin: 1rem 0;
  }
</style>
