package zendesk

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// View represents the views within Zendesk where tickets are grouped and
// observed from.
//  https://developer.zendesk.com/rest_api/docs/support/views#json-format
type View struct {
	ID          int64  `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Active      bool   `json:"active,omitempty"`
	Restriction struct {
		Type string `json:"type,omitempty"`
		ID   int64  `json:"id"`
	} `json:"restriction,omitempty"`
	Position  int64 `json:"position,omitempty"`
	Execution struct {
		GroupBy    string `json:"group_by,omitempty"`
		SortBy     string `json:"sort_by,omitempty"`
		GroupOrder string `json:"group_order,omitempty"`
		SortOrder  string `json:"sort_order,omitempty"`
		Columns    []struct {
			ID    int64  `json:"id,omitempty"`
			Title string `json:"title,omitempty"`
		} `json:"columns,omitempty"`
		Group struct {
			ID    int64  `json:"id,omitempty"`
			Title string `json:"title,omitempty"`
		} `json:"group,omitempty"`
		Sort struct {
			ID    int64  `json:",omitempty"`
			Title string `json:",omitempty"`
		} `json:",omitempty"`
	} `json:"execution,omitempty"`
	Conditions struct {
		All []struct {
			Field    string `json:"field,omitempty"`
			Operator string `json:"operator,omitempty"`
			Value    string `json:"value,omitempty"`
		} `json:"all,omitempty"`
		Any []struct {
			Field    string `json:"field,omitempty"`
			Operator string `json:"operator,omitempty"`
			Value    string `json:"value,omitempty"`
		} `json:"any,omitempty"`
	} `json:"conditions,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type ViewAPI interface {
	GetViews(ctx context.Context) ([]View, Page, error)
	GetView(ctx context.Context, id int64) (View, error)
	CreateView(ctx context.Context, view View) (View, error)
}

func (z *Client) GetViews(ctx context.Context) ([]View, Page, error) {
	var data struct {
		Views []View `json:"views"`
		Page
	}
	u, err := addOptions("/views.json", nil)
	if err != nil {
		return nil, Page{}, err
	}
	body, err := z.get(ctx, u)
	if err != nil {
		return nil, Page{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, Page{}, err
	}
	return data.Views, data.Page, nil
}

func (z *Client) GetActiveViews(ctx context.Context) ([]View, Page, error) {
	var data struct {
		Views []View `json:"views"`
		Page
	}
	u, err := addOptions("/views/active.json", nil)
	if err != nil {
		return nil, Page{}, err
	}
	body, err := z.get(ctx, u)
	if err != nil {
		return nil, Page{}, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, Page{}, err
	}
	return data.Views, data.Page, nil
}

func (z *Client) GetView(ctx context.Context, viewID int64) (View, error) {
	var result struct {
		View View `json:"view"`
	}

	var builder includeBuilder

	u, err := builder.path(fmt.Sprintf("/views/%d.json", viewID))

	if err != nil {
		return View{}, err
	}

	body, err := z.get(ctx, u)
	if err != nil {
		return View{}, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return View{}, err
	}
	return result.View, nil
}

func (z *Client) CreateView(ctx context.Context, view View) (View, error) {
	var data, result struct {
		View View `json:"View"`
	}
	data.View = view

	body, err := z.post(ctx, "/views.json", data)
	if err != nil {
		return View{}, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return View{}, err
	}
	return result.View, nil
}

func (z *Client) UpdateView(ctx context.Context, viewID int64, view View) (View, error) {
	var data, result struct {
		View View `json:"View"`
	}
	data.View = view
	var builder includeBuilder

	u, err := builder.path(fmt.Sprintf("/views/%d.json", viewID))

	body, err := z.put(ctx, u, data)
	if err != nil {
		return View{}, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return View{}, err
	}
	return result.View, nil
}
